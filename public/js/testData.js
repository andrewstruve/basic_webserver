// function : submitData
// request type : POST
// header type : "application/x-www-form-urlencoded"
// parameters: firstName String
//             lastName String
// return Type - returns json object with response and Data collected
// purpose: Submits First and Last name to web server 
function submitData()
{
    var firstName = document.getElementById("firstName").value;
    var lastName = document.getElementById("lastName").value;
    var postString = "firstName=" + firstName +"&"+"lastName=" + lastName;
    xmlhttp = new XMLHttpRequest();
    xmlhttp.onreadystatechange = function() {
    if (this.readyState == 4 && this.status == 200) 
        {
            jsonResponse = this.responseText;
            //JSONResponseObj = JSON.parse(jsonResponse);
            document.getElementById("submitResponse").innerHTML = jsonResponse;        
        }
                    
    };

    xmlhttp.open("POST","/formInput");
    xmlhttp.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    xmlhttp.send(postString);	
}
// function : submitJSON
// request type : POST
// header type : "application/json"
// parameters: JSON parameters
// return Type - returns json object with response and Data collected
// purpose: Submits a json string 
function submitJSON()
{
    var parameter1 = document.getElementById("parameter1").value;
    var parameter2 = document.getElementById("parameter2").value;
    var jsonString = '{ "parameter1" : \"' + parameter1 +'\" , "parameter2": \"' + parameter2 + '\"}'
    console.log(jsonString);
    var jsonObj = JSON.parse(jsonString);
    console.log(jsonObj);
    xmlhttp = new XMLHttpRequest();
    xmlhttp.onreadystatechange = function() {
    if (this.readyState == 4 && this.status == 200) 
        {
            jsonResponse = this.responseText;
            document.getElementById("submitJSONResponse").innerHTML = jsonResponse;        
        }
                    
    };

    xmlhttp.open("POST","/jsonInput");
    xmlhttp.setRequestHeader("Content-type", "application/json");
    xmlhttp.send(jsonObj);	
}
// function : getRandomNumber
// request type : GET
// parameters: None
// return Type - Integer
// purpose: gets a random number
function getRandomNumber()
{
    xmlhttp = new XMLHttpRequest();
    xmlhttp.onreadystatechange = function() {
    if (this.readyState == 4 && this.status == 200) 
        {
            randomNumber = this.responseText;   
            document.getElementById("randomNumber").innerHTML = randomNumber;
        }         
    };
    xmlhttp.open("GET","/getRandomNumber");
    xmlhttp.send();	   
}
// function : getRandomNumberWithParameter
// request type : GET
// parameters: maxNum
// return Type - Integer
// purpose: gets a random number
function getRandomNumberWithParameter()
{
    var parameter1 = document.getElementById("maxNumber").value;
    xmlhttp = new XMLHttpRequest();
    xmlhttp.onreadystatechange = function() {
    if (this.readyState == 4 && this.status == 200) 
        {
            randomNumber = this.responseText;   
            document.getElementById("randomNumberWithParameter").innerHTML = randomNumber;
        }         
    };
    xmlhttp.open("GET","/getRandomNumberParams?maxNum="+parameter1);
    xmlhttp.send();	   
}