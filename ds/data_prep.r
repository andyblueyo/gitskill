library(dplyr)
library(jsonlite)
library(tidyr)

# repo_data <- fromJSON("./data/repos.json")
# 
# user_data <- fromJSON("./data/users.json")
user_data <- flatten(user_data, recursive = TRUE)
View(user_data)
repo_data <- flatten(repo_data, recursive = TRUE)
View(repo_data)
write.csv(x = user_data, file = "./data/users.csv")
write.csv(x = repo_data, file = "./data/repos.csv")