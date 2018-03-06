library(dplyr)
library(jsonlite)
library(tidyr)
library(purrr)

repo_data <- fromJSON("data/data/repos.json")
user_data <- fromJSON("data/data/users.json")

user_data_flat <- flatten(head(user_data, 100), recursive = TRUE) #%>% sample_n(20, replace = FALSE)
repo_data_flat <- flatten(head(repo_data, 100), recursive = TRUE) #%>% sample_n(20, replace = FALSE)

# View(user_data)
# View(repo_data)

df_repo <- data.frame(repo_data_flat[, c( "ownerName", "name", "languages", "ownerType")]) 
df_repo <- df_repo %>% unnest(languages) 
df_repo <- df_repo %>% group_by(ownerName, name1) %>% summarise_at(c("lines"), sum)

df_user <- user_data_flat %>% filter(userType == "User") %>% select(ownerName, userType, publicRepos, orgs) 
df_user <- df_user %>% filter(!map_lgl(orgs, is.null)) %>% unnest() %>% right_join(select(df_user, ownerName))

