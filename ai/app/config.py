# default config
class BaseConfig(object):
    DEBUG = False
    SQLALCHEMY_DATABASE_URI = 'postgresql://onelocation:EUt82sXUYhicJZMI8ZPs@localhost:5432/onelocation'
    SQLALCHEMY_TRACK_MODIFICATIONS = False
    JSON_AS_ASCII = False
    ENDPOINT = "https://southcentralus.api.cognitive.microsoft.com/customvision/v3.0/Prediction/0c667159-2b5c-4449-8aa5-a670fb31edd8/classify/iterations/Iteration2/image"
    PREDICTION_KEY = "284caa38d879463b90d0871031c19958"


class DevelopmentConfig(BaseConfig):
    DEBUG = True


class ProductionConfig(BaseConfig):
    DEBUG = False
