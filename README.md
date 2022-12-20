# todo list

This is a sample application for the IBM Cloud Solution Tutorial [Accelerate a dynamic website using Dynamic Content Acceleration with IBM CDN](https://cloud.ibm.com/docs/solution-tutorials?topic=solution-tutorials-dynamic-content-cdn).


# Build the application locally

1. Install [GO](https://go.dev/). 
  
2. Get the application code. 
   ```
   git clone https://github.com/IBM-Cloud/cdn-with-cda-todolist.git
   ```

3. Run go mod tidy 
   ```
   go mod tidy
   ```

4. From your local to-do application directory, for example, `$GOPATH/src/github.com/IBM-Cloud/cdn-with-cda-todolist`, run the application locally:
   ```bash 
   go build main.go
   ./main
   ```

5. After the application starts, navigate to the URL `http://localhost:8080/` from your browser.

