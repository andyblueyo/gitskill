library(dplyr)
library(jsonlite)
library(tidyr)
library(purrr)
library(tibble)

repo_data <- fromJSON("data/data/repos.json")
user_data <- fromJSON("data/data/users.json")

# Small data frames, only first 100 rows
repo_small <- head(repo_data, 100)
user_small <- head(user_data, 100)

# Tidy repo data, unlist langauages and lines of code
df_repo <- repo_data %>% select(ownerName, name, languages, ownerType) %>% 
  filter(!map_lgl(languages, is.null)) %>% unnest() %>% right_join(select(repo_data, ownerName))
names(df_repo)[names(df_repo) == 'name1'] <- 'languages' #re-name column name
df_repo <- df_repo %>% group_by(ownerName, languages) %>% summarise_at(c("lines"), sum)
df_repo_spread <- df_repo %>% spread(languages, lines)

# Tidy data, unlist orgs, remove organizations
df_user <- user_data %>% select(githubUsername, userType, publicRepos, orgs) %>% filter(userType == "User") 
df_user <- df_user %>% filter(!map_lgl(orgs, is.null)) %>% unnest() %>% right_join(select(df_user, githubUsername))

df_user_repo <-  repo_data %>% select(ownerType, ownerName, fullName, name, languages) %>% mutate(languages = as.character(languages)) %>% unnest(languages) 

  group_by(ownerName, name1)%>% summarise_at(c("lines"), sum)