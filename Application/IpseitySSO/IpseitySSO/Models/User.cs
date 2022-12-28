using AspNetCore.Identity.MongoDbCore.Models;
using MongoDbGenericRepository.Attributes;

namespace IpseitySSO.Models
{
    [CollectionName("Users")]
    public class User : MongoIdentityRole<Guid>
    {
        ////public string Name { get; set; } = string.Empty;
        //public UserAdditionalEmails? AdditionalEmails { get; set; }
        //public UserAvatar[] Avatars { get; set; } = null!;

    }

    public class UserAdditionalEmails
    {
        public string Email { get; set; } = string.Empty;
        public string NormalizedEmail { get; set; } = string.Empty;
        public bool EmailConfirmed { get; set; }
    }

    public class UserAvatar
    {
        public Boolean Default { get; set; }

        public string Image { get; set; } = string.Empty;
    }
}
