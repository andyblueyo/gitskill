"use strict";

var gulp = require("gulp");
var sass = require("gulp-sass");

gulp.task("sass", function() {
    return gulp
        .src("./scss/style.scss")
        .pipe(sass({ outputStyle: "compressed" }).on("error", sass.logError))
        .pipe(gulp.dest("./"));
});

gulp.task("watch", function() {
    gulp.watch("./scss/**/*.scss", ["sass"]);
});

gulp.task("default", ["sass", "watch"]);
