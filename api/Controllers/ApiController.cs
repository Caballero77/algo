using System.Collections.Generic;
using System.Threading.Tasks;
using api.Models;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Configuration;
using MongoDB.Bson;
using MongoDB.Driver;

namespace api.Controllers
{
    [Route("api")]
    public class ApiController : ControllerBase
    {
        private readonly IConfiguration configuration;

        private readonly IMongoCollection<SortingModel> collection;

        public ApiController(IConfiguration configuration)
        {
            this.configuration = configuration;

            var mongoClient = new MongoClient(configuration.GetValue("mongo", "mongodb://localhost:27017"));
            var db = mongoClient.GetDatabase(configuration.GetValue("database", "sorting"));
            this.collection = db.GetCollection<SortingModel>(configuration.GetValue("collection", "results"));
        }

        [HttpPost]
        public async Task<int> AddResult([FromBody] SortingModel model)
        {
            await this.collection.InsertOneAsync(model);

            return model.Id;
        }

        [HttpGet]
        public Task<List<SortingModel>> GetModels([FromQuery] int take, [FromQuery] int skip, [FromQuery] string type)
            => this.collection.Find(d => d.SortingType == type).Skip(skip).Limit(take).ToListAsync();
    }
}