from fastapi import FastAPI, status, Request
from api.routes import api_router
from fastapi.exceptions import RequestValidationError
from fastapi.responses import JSONResponse


def include_router(app):
    app.include_router(api_router)


def start_services():
    app = FastAPI(title="Receipt Processor")

    include_router(app)

    return app


app = start_services()


@app.exception_handler(RequestValidationError)
async def validation_exception_handler(request: Request, exc: RequestValidationError):
    return JSONResponse(
        status_code=status.HTTP_400_BAD_REQUEST,
        content={"detail": exc.errors()},
    )
