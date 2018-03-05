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
df_repo <- data.frame(repo_data_flat[, c("name", "languages", "ownerType", "ownerName")]) 

z <- 1
for(language in df_repo$languages){
    if(length(language)>0){
        
        lines <- language[,"lines"]
        lang <- language[,"name"]
        if(length(lang)>0){
            
            for(i in 1:length(lang)){
                print(z)
                
                df_repo[z, lang[i]] <- lines[i]
            }
        }
    }
    
    z <- z + 1
}

drops <- c("languages")
df_repo2 <- df_repo[, !(names(df_repo) %in% drops)] #%>% apply(2,as.character)
df_repo2[is.na(df_repo2)] <- 0

df_repo3 <- df_repo2 %>% 
    group_by(ownerName) %>% 
    filter() %>% 
    summarize_all(funs(sum(as.numeric(.))))

View(df_repo)
