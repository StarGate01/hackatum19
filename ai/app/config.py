# default config
class BaseConfig(object):
    DEBUG = False
    SQLALCHEMY_DATABASE_URI = 'postgresql://onelocation:EUt82sXUYhicJZMI8ZPs@localhost:5432/onelocation'
    SQLALCHEMY_TRACK_MODIFICATIONS = False
    JSON_AS_ASCII = False


class DevelopmentConfig(BaseConfig):
    DEBUG = True


class ProductionConfig(BaseConfig):
    DEBUG = False
