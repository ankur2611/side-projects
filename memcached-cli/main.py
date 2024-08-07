import memcache
import click


DEFAULT_HOST = 'localhost'
DEFAULT_PORT = 11211


def connect(host, port):
    return memcache.Client([f'{host}:{port}'])


@click.group()
@click.option('-host', default=DEFAULT_HOST, help='Memcached server hostname')
@click.option('-port', default=DEFAULT_PORT, type=int, help='Memcached server port')
@click.pass_context
def cli(ctx, host, port):
    '''A simple Memcached CLI client'''
    ctx.ensure_object(dict)
    ctx.obj['CLIENT'] = connect(host, port)


@cli.command()
@click.argument('key')
@click.argument('value')
@click.option('-ttl', default=0, help='Expiration time for the data')
@click.pass_context
def set(ctx, key, value, ttl):
    '''Set the value of a key'''
    client = ctx.obj['CLIENT']
    click.echo(f'STORED') if client.set(key, value, time=ttl) else click.echo(f'NOT_STORED')


@cli.command()
@click.argument('key')
@click.argument('value')
@click.option('-ttl', default=0, help='Expiration time for the data')
@click.pass_context
def add(ctx, key, value, ttl):
    '''Add the value of a non-existing/new key'''
    client = ctx.obj['CLIENT']
    click.echo(f'STORED') if client.add(key, value, time=ttl) else click.echo(f'NOT_STORED')


@cli.command()
@click.argument('key')
@click.argument('value')
@click.option('-ttl', default=0, help='Expiration time for the data')
@click.pass_context
def replace(ctx, key, value, ttl):
    '''Replace existing key with value'''
    client = ctx.obj['CLIENT']
    click.echo(f'STORED') if client.replace(key, value, time=ttl) else click.echo(f'NOT_STORED')


@cli.command()
@click.argument('key')
@click.argument('value')
@click.option('-ttl', default=0, help='Expiration time for the data')
@click.pass_context
def append(ctx, key, value, ttl):
    '''Append the value to the end of the existing key's value'''
    client = ctx.obj['CLIENT']
    click.echo(f'STORED') if client.append(key, value, time=ttl) else click.echo(f'NOT_STORED')


@cli.command()
@click.argument('key')
@click.argument('value')
@click.option('-ttl', default=0, help='Expiration time for the data')
@click.pass_context
def prepend(ctx, key, value, ttl):
    '''Prepend the value to the beginning of the existing key's value'''
    client = ctx.obj['CLIENT']
    click.echo(f'STORED') if client.prepend(key, value, time=ttl) else click.echo(f'NOT_STORED')


@cli.command()
@click.argument('key')
@click.pass_context
def get(ctx, key):
    '''Retrieves a key from the memcache'''
    client = ctx.obj['CLIENT']
    value = client.get(key)
    click.echo(f'{value}\nEND')


@cli.command()
@click.argument('key')
@click.pass_context
def delete(ctx, key):
    '''Deletes a key from the memcache'''
    client = ctx.obj['CLIENT']
    click.echo(f'DELETED') if client.delete(key) else click.echo(f'NOT_DELETED')


@cli.command()
@click.argument('key')
@click.argument('delta')
@click.pass_context
def increment(ctx, key, delta):
    '''Increment value for existing key by delta'''
    client = ctx.obj['CLIENT']
    click.echo(f'INCREMENTED') if client.incr(key, delta) else click.echo(f'NOT_INCREMENTED')


@cli.command()
@click.argument('key')
@click.argument('delta')
@click.pass_context
def decrement(ctx, key, delta):
    '''Decrement value for existing key by delta'''
    client = ctx.obj['CLIENT']
    click.echo(f'DECREMENTED') if client.decr(key, delta) else click.echo(f'NOT_DECREMENTED')


@cli.command()
@click.argument('key')
@click.argument('value')
@click.option('-ttl', default=0, help='Expiration time for the data')
@click.pass_context
def cas(ctx, key, value, ttl):
    '''Sets a key to a given value in the memcache if it hasn't been altered since last fetched'''
    client = ctx.obj['CLIENT']
    click.echo(f'STORED') if client.cas(key, value, time=ttl) else click.echo(f'NOT_STORED')
