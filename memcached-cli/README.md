**Memcached CLI**

A command-line interface (CLI) tool for interacting with a Memcached server.

***Installation***

First setup a memcached server.
Afterwards, to install the package, navigate to the directory containing the `setup.py` file and run:

    pip install .

***Usage***

This CLI tool provides several commands to interact with a Memcached server. Below are the available commands and their usage:

**Set a Key**

Sets the value of a key.

    ccmm set <key> <value> [-ttl <expiration_time>]

**Add a Key**

Adds the value of a non-existing/new key.

    ccmm add <key> <value> [-ttl <expiration_time>]

**Replace a Key**

Replaces the value of an existing key.

    ccmm replace <key> <value> [-ttl <expiration_time>]

**Append to a Key**

Appends the value to the end of the existing key's value.

    ccmm append <key> <value> [-ttl <expiration_time>]

**Prepend to a Key**

Prepends the value to the beginning of the existing key's value.

    ccmm prepend <key> <value> [-ttl <expiration_time>]

**Get a Key**

Retrieves a key from the Memcached server.

    ccmm get <key>

**Delete a Key**

Deletes a key from the Memcached server.

    ccmm delete <key>

**Increment a Key**

Increments the value of an existing key by a specified delta.

    ccmm increment <key> <delta>

**Decrement a Key**

Decrements the value of an existing key by a specified delta.

    ccmm decrement <key> <delta>


**CAS (Check and Set)**

Sets a key to a given value in the Memcached server if it hasn't been altered since it was last fetched.

    ccmm cas <key> <value> [-ttl <expiration_time>]


*Example*
```
# Set a key with a value and a TTL of 60 seconds
ccmm set my_key my_value -ttl 60

# Add a new key with a value
ccmm add new_key new_value

# Replace an existing key with a new value
ccmm replace existing_key new_value

# Append a value to an existing key
ccmm append existing_key additional_value

# Prepend a value to an existing key
ccmm prepend existing_key additional_value

# Get the value of a key
ccmm get my_key

# Delete a key
ccmm delete my_key

# Increment the value of a key by 5
ccmm increment my_key 5

# Decrement the value of a key by 3
ccmm decrement my_key 3

# Set a key with a value using CAS with a TTL of 60 seconds
ccmm cas my_key my_value -ttl 60
```


**Author**

Ankur Mehta

This `README.md` provides an overview of the CLI tool, installation instructions, usage examples, and other relevant information.
