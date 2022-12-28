using System;
namespace IpseitySSO.Dtos.Responses
{
    public class LoginResponse : Response
    {
        public string AccessToken { get; set; } = string.Empty;
        public string Email { get; set; } = string.Empty;
        public string UserId { get; set; } = string.Empty;
    }
}
