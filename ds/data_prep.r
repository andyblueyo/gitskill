library(dplyr)
library(jsonlite)
library(tidyr)

repo_data <- fromJSON("./data/repos.json")
user_data <- fromJSON("./data/users.json")

user_data_flat <- flatten(user_data, recursive = TRUE) %>% sample_n(20, replace = FALSE)
repo_data_flat <- flatten(repo_data, recursive = TRUE) %>% sample_n(20, replace = FALSE)

# View(user_data)
# View(repo_data)

print(repo_data[2,6])
df <- data.frame(repo_data_flat[, c("name", "languages", "ownerType", "ownerName")]) 

View(df)
