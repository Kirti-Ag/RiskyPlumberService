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



