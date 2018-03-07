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

df_repo <- data.frame(repo_data_flat[, c( "ownerName", "languages", "ownerType", "stars", "forks")]) 

df_user_repo <-  repo_data %>% select(ownerType, ownerName, fullName, name, languages) %>% mutate(languages = as.character(languages)) %>% unnest(languages) 

  group_by(ownerName, name1)%>% summarise_at(c("lines"), sum)

index <- 1
#Loop through language lists
for(language in df_repo$languages){
  
    if(length(language)>0){
      
        #If there are any languages in the list, get the name and the lines
        lines <- language[,"lines"]
        lang <- language[,"name"]
        if(length(lang)>0){
            
            for(i in 1:length(lang)){
                #Add cell in column of the language and index of the person we want
                df_repo[index, lang[i]] <- lines[i]
            }
        }
    }
    
  index <- index + 1
}

# remove languages and set NAs to 0
drops <- c("languages")
df_repo2 <- df_repo[, !(names(df_repo) %in% drops)] #%>% apply(2,as.character)
df_repo2[is.na(df_repo2)] <- 0

# Group by user 
df_repo2 <- df_repo2 %>% 
    group_by(ownerName, ownerType) %>% 
    summarize_all(funs(sum(as.numeric(.))))

#make column names match for merge
colnames(user_data_flat)[1] <- "ownerName"

#Select JUST Users (not orgs) from both dataframes
#df_repo2 <- df_repo2 %>% filter(ownerType == "User")
user_data_df <- data.frame(user_data_flat[, c( "ownerName", "userType","publicRepos", "orgs")]) 
  
#Merge 
merged_dfs <- merge(user_data_df, df_repo2, by = "ownerName", all = TRUE) %>% apply(2,as.character)

write.csv(x = merged_dfs, file = "./data/merged_data.csv")
colnames(df4)
View(df_repo)

