#   go-fibonacci

#### REST API WHICH SENDS BACK THE CORESSPONDING FIBONACCI NUMBER OF ANY GIVEN NUMBER
###### Author Daniyal Ibrahim
###### Date   10.10.2020
###### Go lang 1.15 with gorilla and gin 

The following steps will help understanding what is happening in this chunk of code.

##### STEP 1 Create a public git repository at Github

    mkdir myDirName    #this is the name of your directory
    cd /myDirName
    ssh -T git@github.com # authenticate if you are logged in at your terminal
    git init
    touch readME.md   #this is to create an initial file to push
    git commit -m "enter commit message here"
    git remote add origin git@github.com:YOUR_USERNAME/myDirName.git
    
this will create the repo in github.
    
    curl -u USERNAME:PASSWORD https://api.github.com/user/repos -d '{"name":"myDirName"}' 

if you haven't generated and SSH key for github access then follow these steps, otherwise you're good to push your shit to github.

      eval $(ssh-agent -s)
      ssh-keygen -t rsa -b 4096 -C "email@yourdomain.com" #this should be your github email address

you'll be prompted to a couple of times. Press enter for the first prompt. choose a passphrase for the second prompt, or press enter twice for no passphrase

      ssh-add ~/.ssh/id_rsa   #this is your private key
      cat ~/.ssh/id_rsa.pub   # copy the output of this command. this is your SSH public key
      curl -u USERNAME:PASSWORD https://api.github.com/user/keys -d '{"title":"KEY_NAME", "key":"YOUR_RSA_PUBLIC_KEY_HERE"}'

the value you copied earlier and your keyname.I recommend using a combination of machine name and app (My-laptop (Git CLI)
      
      git push -u origin master
      
##### Following steps

    1. create a new project in Go lang. I prefer Go Land IDE from jet Brains
    2. commit the state to the Git repository, add the TAG "3_NewApi
    3. create a new REST service with the route /helloapi/fibonacci
    4. the service should be able to be called via HTTP GET and take a parameter N (integer, greater than zero)
        can.For each number passed in the parameter, the API calculates the corresponding Fibonacci number.
    5. as a result, the API returns both N and the calculated Fibonacci number. The return formats are XML or
        JSON in question.
    6. committee the state in the git repository, add the TAG "8_Fibonacci
        If >= 12 is passed as a parameter for N, the calculated value should exceptionally not follow the Fibonacci sequence
        but always get the value 100.
    7. committee the state in the git repository, add the TAG "9_PlanningPoker".
    8. create a HTTP client request with a tool of your choice (Fiddler, Postman, CURL, etc) and save both the complete HTTP request and the associated response in a text file.
    9. commit the state to the git repository, add the TAG "12_HttpClient
