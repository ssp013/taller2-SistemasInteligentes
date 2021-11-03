//import User from "./class";
const User = require("./class.js");//commún
const hitesh = new User("Seba","ssp013@gmail.com");
hitesh.courseList.push("math");
console.log(hitesh.getCourselist());