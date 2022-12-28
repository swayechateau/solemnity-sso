
using AspNetCore.Identity.MongoDbCore.Models;
using MongoDbGenericRepository.Attributes;

namespace IpseitySSO.Models
{
    [CollectionName("Users")]
    public class User : MongoIdentityUser<Guid>
    {
        public string Name { get; set; } = string.Empty;
        public IList<UserAdditionalEmails>? AdditionalEmails { get; set; }
    }

    public class UserAdditionalEmails
    {
        public string Email { get; set; } = string.Empty;
        public string NormalizedEmail { get; set; } = string.Empty;
        public bool EmailConfirmed { get; set; }
    }

}