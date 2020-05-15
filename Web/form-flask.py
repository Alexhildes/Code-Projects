from flask import Flask, render_template, request

app = Flask(__name__)

@app.route('/', methods=['GET', 'POST'])
def index():
    print(request.method)
    if request.method = 'POST':
        print(request.form)
    
    return render_template('brew.html')

if __name__ == "__main__":
    app.run(host='0.0.0.0', port=80, debug=True)
