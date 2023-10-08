#!/usr/bin/env python3.11
import pymongo
import json

def load_mongo_data(file_path, db_name, mongo_uri="mongodb://admin:password@localhost:27017/"):
    client = pymongo.MongoClient(mongo_uri)
    db = client[db_name]
    
    
    with open(file_path, 'r') as file:
        resume = json.load(file)
        for key in resume.keys():
            collection = db[key]
            print(f"Loading data into {db_name}.{collection}...")
            collection.insert_one(resume[key])
            print(f"Loaded data from {file_path} into {db_name}.{collection}.")

if __name__ == "__main__":
    load_mongo_data("./resume/resume.json", "resume")