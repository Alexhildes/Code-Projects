from flask import Flask, render_template, request


app = Flask(__name__)


@app.route("/", methods=['GET', 'POST'])
def index():
    print(request.method)
    if request.method == 'POST':
        print(request.form)
        #get method looks for corresponding "name" in HTML == "value"
        
            return render_template("test.html")
    elif request.method == 'GET':
        # return render_template("index.html")
        print("No Post Back Call")
    return render_template("test.html")

if __name__ == "__main__":
    app.run(host='0.0.0.0', port=80, debug=True)

