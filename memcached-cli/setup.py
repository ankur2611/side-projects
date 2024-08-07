from setuptools import setup


setup(
    author='Ankur Mehta',
    name='memcached_cli',
    description='A simple command line interface for memcached',
    version='0.1',
    py_modules=['main'],
    install_requires=[
        'click',
        'python-memcached',
    ],
    entry_points='''
        [console_scripts]
        ccmm=main:cli
    ''',
)
