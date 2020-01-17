using System.Collections.Generic;
using MongoDB.Bson.Serialization.Attributes;

namespace api.Models
{
    public class SortingModel
    {
        [BsonId]
        public int Id { get; set; }

        public string SortingType { get; set; }

        public IEnumerable<int> Array { get; set; }

        public IEnumerable<int> Result { get; set; }
    }
}