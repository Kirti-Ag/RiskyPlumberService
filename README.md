From the command line(terminal) in the directory containing main.go, run the code.
<pre><code>go get .
</code></pre>
<pre><code>go run .
</code></pre>

You will see:
<pre><code>Listening and serving HTTP on :8080
</code></pre>

From a new command line window(terminal), use curl to make a request to your running web service.

To get the list of all risks:
<pre><code>curl http://localhost:8080/v1/risks
</code></pre>

To add a risk:
<pre><code>curl http://localhost:8080/v1/risks \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"state": "open", "title": "Risk1","description": "First Risk"}'
</code></pre>

To get risk by id:
<pre><code>curl http://localhost:8080/v1/risks/<id>
</code></pre>
Here replace <code>id</code> with risk UUID.


Here is a sample run of the web service.
---
From the command line(terminal) in the directory containing main.go, run the code.
![alt text](https://github.com/Kirti-Ag/RiskyPlumberService/blob/main/SampleRun1.png)

From a new command line window(terminal), use curl to make a request to your running web service.
Response to first <code>GET</code> request to <code>/v1/risks</code> is empty list:
![alt text](https://github.com/Kirti-Ag/RiskyPlumberService/blob/main/img1.png)

Creating a new risk, the command displays headers and JSON for the added risk.
![alt text](https://github.com/Kirti-Ag/RiskyPlumberService/blob/main/img2.png)

Creating a new risk with invalid <code>state</code attribute, the command displays headers with 500 Internal Server Error code and error message.
![alt text](https://github.com/Kirti-Ag/RiskyPlumberService/blob/main/img3.png)

Creating a new risk, the command displays headers and JSON for the added risk. Following up with getting all risks and getting risks by <code>id</code> of the second risk.
![alt text](https://github.com/Kirti-Ag/RiskyPlumberService/blob/main/img4.png)

Thank you for following through!


