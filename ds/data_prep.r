library(dplyr)
library(jsonlite)
library(tidyr)

repo_data <- fromJSON("./data/repos.json")
user_data <- fromJSON("./data/users.json")

user_data_flat <- flatten(user_data, recursive = TRUE) #%>% sample_n(20, replace = FALSE)
repo_data_flat <- flatten(repo_data, recursive = TRUE) #%>% sample_n(20, replace = FALSE)

# View(user_data)
# View(repo_data)

df_repo <- data.frame(repo_data_flat[, c( "ownerName", "languages", "ownerType", "stars", "forks")]) 


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
