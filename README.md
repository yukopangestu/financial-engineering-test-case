# financial-engineering-test-case
Hi engineering team that review this codes. Let me introduce myself, my name is Yuko Pangestu, one of the candidate for your company. 

So, for the challenge i have chosen the third question regarding borrowers and investor in loan.

# System Requirements
- Docker (needed for multiple environment build situation)

If you don't have any docker environment, you can install
1. Go version 1.25.4
2. MySQL version 8.0.21

Should be enough but again, what I think is quickest and less painful is using the docker itself, just run `docker compose up` and you just good to go using your own API request software like Postman. Also there is a file that you can import to Postman to test the API, you just need to adjust the loan_id environment each request.

# API Request Explanation

In this system, I imagine there will be 4 services:
1. Borrower Service
2. Investor Service
3. Employee Service
4. Loan Service

Why 4? Because I think it's a good idea to have a service that deals with the loan itself, and the other 3 services are just for the data management. That way, the loan service will be the only one that deals with the business logic of the loan.

In this case, I just provide the creation API only, I think that's enough for the challenge because what I built is that loan service is dependant to the remaining services to be able to run. So, in order to run, you need to add some data to borrower, investor, and employee tables, then you can add the data based on the id provided on each table.

# ENV

The env provided here is regarding 2 things
1. MySQL -> I'm using the docker image provided by the composer
2. SMTP -> I'm using the mailtrap.io service, you can use your own SMTP service if you want

I think that's all, feel free to ask me if you have any questions. I'm more open to discuss how to improve this system also.

Thank you,<br>
Yuko Pangestu