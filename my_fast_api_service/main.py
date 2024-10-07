from fastapi import FastAPI

app = FastAPI(title="Receipt Processor")

@app.get("/")
def welome():
    return {"msg": "Hello, world!"}

