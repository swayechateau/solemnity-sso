using System;
using MongoDB.Driver.Core.Configuration;

namespace IpseitySSO
{
	public class DatabaseSettings
	{
        public string ConnectionString { get; set; } = null!;

        public string DatabaseName { get; set; } = null!;

        public string UsersCollectionName { get; set; } = null!;

    }
}

