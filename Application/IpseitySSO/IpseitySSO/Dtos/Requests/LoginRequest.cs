﻿using System;
using System.ComponentModel.DataAnnotations;

namespace IpseitySSO.Dtos.Requests
{
    public class LoginRequest
    {
        [Required(ErrorMessage = "Email is required"), EmailAddress]
        public string Email { get; set; } = string.Empty;
        [Required(ErrorMessage = "Password is required"), DataType(DataType.Password)]
        public string Password { get; set; } = string.Empty;
	}
}

