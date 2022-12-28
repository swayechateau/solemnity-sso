using AspNetCore.Identity.MongoDbCore.Models;
using MongoDbGenericRepository.Attributes;

namespace IpseitySSO.Models
{
    [CollectionName("Roles")]
    public class Role : MongoIdentityRole<Guid>
	{

	}
}

