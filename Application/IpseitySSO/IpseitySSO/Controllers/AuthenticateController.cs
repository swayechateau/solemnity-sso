using System;
using System.Collections.Generic;
using System.Data;
using System.IdentityModel.Tokens.Jwt;
using System.Linq;
using System.Net;
using System.Security.Claims;
using System.Text;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Identity;
using Microsoft.AspNetCore.Mvc;
using Microsoft.IdentityModel.Tokens;
using IpseitySSO.Dtos.Responses;
using IpseitySSO.Dtos.Requests;
using IpseitySSO.Models;

namespace IpseitySSO.Controllers
{
    [ApiController]
    [Route("api/[controller]")]
    public class AuthenticateController : Controller
    {
        private readonly UserManager<User> _userManager;
        private readonly RoleManager<Role> _roleManager;
        private readonly IConfiguration _configuration;

        public AuthenticateController(
           UserManager<User> userManager,
           RoleManager<Role> roleManager,
           IConfiguration configuration)
        {
            _userManager = userManager;
            _roleManager = roleManager;
            _configuration = configuration;
        }
        [HttpPost]
        [Route("roles/add")]
        public async Task<IActionResult> CreateRole([FromBody] RoleRequest request) =>
            await CreateRoleAsync(request);

        [HttpPost]
        [Route("register")]
        public async Task<IActionResult> Register([FromBody] RegisterRequest request) =>
            await RegisterAsync(request);

        //[HttpPost]
        //[Route("register-admin")]
        //public async Task<IActionResult> RegisterAdmin([FromBody] RegisterRequest request) =>
        //    await RegisterAdminAsync(request);

        [HttpPost]
        [Route("login")]
        public async Task<IActionResult> Login([FromBody] LoginRequest request) =>
            await LoginAsync(request);


        //private async Task<IActionResult> RegisterAdminAsync(RegisterRequest model)
        //{
        //    var userExists = await _userManager.FindByNameAsync(model.Username);
        //    if (userExists != null)
        //        return StatusCode(StatusCodes.Status500InternalServerError, new Response { Status = "Error", Message = "User already exists!" });

        //    User user = new()
        //    {
        //        Email = model.Email,
        //        SecurityStamp = Guid.NewGuid().ToString(),
        //        UserName = model.Username
        //    };
        //    var result = await _userManager.CreateAsync(user, model.Password);
        //    if (!result.Succeeded)
        //        return StatusCode(StatusCodes.Status500InternalServerError, new Response { Status = "Error", Message = "User creation failed! Please check user details and try again." });

        //    if (!await _roleManager.RoleExistsAsync(UserRoles.Admin))
        //        await _roleManager.CreateAsync(new IdentityRole(UserRoles.Admin));
        //    if (!await _roleManager.RoleExistsAsync(UserRoles.User))
        //        await _roleManager.CreateAsync(new IdentityRole(UserRoles.User));

        //    if (await _roleManager.RoleExistsAsync(UserRoles.Admin))
        //    {
        //        await _userManager.AddToRoleAsync(user, UserRoles.Admin);
        //    }
        //    if (await _roleManager.RoleExistsAsync(UserRoles.Admin))
        //    {
        //        await _userManager.AddToRoleAsync(user, UserRoles.User);
        //    }
        //    return Ok(new Response { Status = "Success", Message = "User created successfully!" });
        //}

        private async Task<IActionResult> CreateRoleAsync(RoleRequest request)
        {
            var appRole = new Role { Name = request.Role };
            var createRole = await _roleManager.CreateAsync(appRole);

            return Ok(new { message = "role created succesfully" });
        }

        private async Task<IActionResult> RegisterAsync(RegisterRequest request)
        {
            try
            {
                var userExists = await _userManager.FindByEmailAsync(request.Email);
                if (userExists != null)
                    return StatusCode(StatusCodes.Status500InternalServerError, new Response
                    {
                        Status = "Error",
                        Message = "User already exists"
                    });
                //if we get here, no user with this email.
                userExists = new User
                {
                    Name = request.FullName,
                    Email = request.Email,
                    ConcurrencyStamp = Guid.NewGuid().ToString(),
                    UserName = request.Username
                };
                var createUserResult = await _userManager.CreateAsync(userExists, request.Password);
                if (!createUserResult.Succeeded)
                    return StatusCode(StatusCodes.Status500InternalServerError, new Response
                    {
                        Status = "Error",
                        Message = $"Create user failed {createUserResult?.Errors?.First()?.Description}"
                    });
                var addUserToRoleResult = await _userManager.AddToRoleAsync(userExists, "USER");
                if (!addUserToRoleResult.Succeeded)
                    return StatusCode(StatusCodes.Status500InternalServerError, new Response
                    {
                        Status = "Error",
                        Message = $"Create user succeeded but could not add user to role {addUserToRoleResult?.Errors?.First()?.Description}"
                    });

                return StatusCode(StatusCodes.Status201Created, new RegisterResponse
                {
                    Status = "Success",
                    Message = "User registered successfully"
                });
            }
            catch (Exception ex)
            {
                Console.WriteLine(ex.Message);
                return StatusCode(StatusCodes.Status500InternalServerError, new Response
                {
                    Status = "Error",
                    Message = ex.Message
                });
            }
        }

        private async Task<IActionResult> LoginAsync(LoginRequest request)
        {

            try
            {
                var user = await _userManager.FindByEmailAsync(request.Email);
                var validUser = await CheckUser(user, request.Password);
                if (!validUser)
                    return Unauthorized();

                var token = GetToken(LoginCaims(user));
                return StatusCode(StatusCodes.Status200OK, new LoginResponse
                {
                    AccessToken = new JwtSecurityTokenHandler().WriteToken(token),
                    //Expiration = token.ValidTo,
                    Email = user?.Email,
                    Status = "Success",
                    UserId = user?.Id.ToString()
                });
            }
            catch (Exception ex)
            {
                Console.WriteLine(ex.Message);
                return StatusCode(StatusCodes.Status500InternalServerError, new LoginResponse
                {
                    Status = "Error",
                    Message = ex.Message
                });
            }

        }

        private async Task<bool> CheckUser(User user, string password) =>
            user != null && await _userManager.CheckPasswordAsync(user, password);

        private JwtSecurityToken GetToken(List<Claim> authClaims)
        {
            var authSigningKey = new SymmetricSecurityKey(Encoding.UTF8.GetBytes(_configuration["JWT:Secret"]));

            var token = new JwtSecurityToken(
                issuer: _configuration["JWT:ValidIssuer"],
                audience: _configuration["JWT:ValidAudience"],
                expires: DateTime.Now.AddHours(3),
                claims: authClaims,
                signingCredentials: new SigningCredentials(authSigningKey, SecurityAlgorithms.HmacSha256)
                );

            return token;
        }

        private List<Claim> LoginCaims(User user)
        {
            var claims = new List<Claim>
                {
                    new Claim(JwtRegisteredClaimNames.Sub, user.Id.ToString()),
                    new Claim(ClaimTypes.Name, user.UserName),
                    new Claim(JwtRegisteredClaimNames.Jti, Guid.NewGuid().ToString()),
                    new Claim(ClaimTypes.NameIdentifier, user.Id.ToString())

                };
            var roles = await _userManager.GetRolesAsync(user);
            var roleClaims = roles.Select(x => new Claim(ClaimTypes.Role, x));
            claims.AddRange(roleClaims);
            return claims;
        }
        //    public async Task<IActionResult> Login([FromBody] LoginRequest model)
        //    {
        //        var user = await _userManager.FindByNameAsync(model.Username);
        //        if (user != null && await _userManager.CheckPasswordAsync(user, model.Password))
        //        {
        //            var userRoles = await _userManager.GetRolesAsync(user);

        //            var authClaims = new List<Claim>
        //            {
        //                new Claim(ClaimTypes.Name, user.UserName),
        //                new Claim(JwtRegisteredClaimNames.Jti, Guid.NewGuid().ToString()),
        //            };

        //            foreach (var userRole in userRoles)
        //            {
        //                authClaims.Add(new Claim(ClaimTypes.Role, userRole));
        //            }

        //            var token = GetToken(authClaims);

        //            return Ok(new
        //            {
        //                token = new JwtSecurityTokenHandler().WriteToken(token),
        //                expiration = token.ValidTo
        //            });
        //        }
        //        return Unauthorized();
        //    }
        //}

    }
}

