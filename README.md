# todo list

This is a sample application for the IBM Cloud Solution Tutorial [Accelerate a dynamic website using Dynamic Content Acceleration with IBM CDN](https://cloud.ibm.com/docs/solution-tutorials?topic=solution-tutorials-dynamic-content-cdn).


# Build the application locally

If you have already prepared the [GO](https://golang.org/) basis and the [Beego](https://beego.me/docs/intro/) framework for application development, you can run the application locally. Otherwise, move on to the next section [Customize the application to include a test object](#customize-test-object).

1. Install [GO](https://go.dev/). 
  
2. Install the Bee tool. 
   ```bash
   go get github.com/beego/bee
   ```
  
3. Install Beego.
   ```
   go get github.com/beego/beego
   ```

   ```
   export GO111MODULE=on
   ```   
	
4. Get the application code. 
   ```
   git clone https://github.com/IBM-Cloud/cdn-with-cda-todolist.git
   ```

5. From your local to-do application directory, for example, `$GOPATH/src/github.com/IBM-Cloud/cdn-with-cda-todolist`, run the application locally:
   ```bash 
   go build main.go
   ./main
   ```
6. After the application starts, navigate to the URL `http://localhost:8080/` from your browser.

