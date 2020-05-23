# TESTING INDIVIDUAL COMPONENTS OF THE PIPELINE

## Blockchain Implementation in GoLang For Placement System

## The Consensus Algorithm implemented in Blockchain System is a combination of Proof Of Work and Proof Of Elapsed Time


### Run `go run main.go` to Start the Server and listen on localhost:5000

### Usage :


#### To Print Usage
####    Make POST request to `/usage`

#### To Print the BlockChain
####    Make POST request to `/print`

#### Manual Pipeline - 

#### To Create a New BlockChain    
####    Make POST request to `/createBlockChain`

#### To Add a New Student
####    Make POST request to `/student` with body -
```json
{
    "Usn": "1MS16CS034",
    "Branch": "CSE",
    "Name": "Gaurav",
    "Gender": "Male",
    "Dob": "30-10-1998",
    "Cgpa": "9",
    "Perc10th": "90",
    "Perc12th": "90",
    "Backlog": false,
    "Email": "gauravkarkal@gmail.com",
    "Mobile": "8867454545",
    "Staroffer": true
}
```


#### To Add a new Company    
####    Make POST request to `/company` with body -
```json
{
    "company": "GE"
}
```

#### To Send Email to Eligible Students based on Eligibility Criteria
####    Make POST request to `/send` with body -
```json
{
	"company" : "JPMC",
	"backlog" : "",
	"starOffer" : "",
	"branch" : ["CSE","ISE"],
	"gender" : "",
	"cgpaCond" : "GreaterThan",
	"cgpa" : "2",
	"perc10thCond" : "GreaterThan",
	"perc10th" : "10",
	"perc12thCond" : "GreaterThan",
	"perc12th" : "10"
}
```
#### To Handle Request and Initiate Creation of Request Block
####    Make GET request to `/handlerequest` with Query Params -
```json
Key :   Value

approval: true
company: JPMC
name: 1MS16CS034

```

#### To Run Verification by Academic Department
####    Make POST request to `/verify-AcademicDept` with body -
```json
{
	"name":"1MS16CS034",
    "company": "JPMC"
  
}
```

#### To Run Verification by Placement Department
####    Make POST request to `/verify-PlacementDept` with body -
```json
{
	"name":"1MS16CS034",
    "company": "JPMC"
  
}
```

#### To Retrieve the data for the Company
####    Make POST request to `/companyRetrieveData` with body -
```json
{
	"name":"1MS16CS034",
    "company": "JPMC"
  
}
```
#### End of Pipeline

#### Test Direct Request to Student
####    Make POST request to `/request-student` with body -
```json
{
	"name":"1MS16CS034",
    "company": "JPMC"
  
}
```



