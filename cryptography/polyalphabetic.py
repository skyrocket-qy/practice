

def polyalphabetic_encrypt(input: str, encrypt_map1: dict, encrypt_map2: dict, repeat_pattern: str) -> str:
    """
    Encrypts a string using a polyalphabetic cipher.

    Args:
        input: The string to encrypt.
        encrypt_map1: The mapping of odd-positioned characters to encrypted characters.
        encrypt_map2: The mapping of even-positioned characters to encrypted characters.
        repeat_pattern: The pattern to repeat the encrypt maps.

    Returns:
        The encrypted string.
    """

    # Check that the repeat pattern is valid.

    if not all(x in "12" for x in repeat_pattern):
        raise ValueError("Invalid repeat pattern")

    encrypted_string = ""
    i = 0
    encrypt_map = encrypt_map1
    for character in input:
        if i % len(repeat_pattern) == 0:
            encrypt_map = encrypt_map1 if i % 2 == 0 else encrypt_map2
        i += 1
        encrypted_character = encrypt_map.get(character)
        if encrypted_character is None:
            encrypted_character = character
        encrypted_string += encrypted_character

    return encrypted_string

if  __name__ == "__main__":
    encrypt_map1 = {
        "a": "b",
        "b": "c",
        "c": "d",
        "d": "e",
        "e": "f",
        "f": "g",
        "g": "h",
        "h": "i",
        "i": "j",
        "j": "k",
        "k": "l",
        "l": "m",
        "m": "n",
        "n": "o",
        "o": "p",
        "p": "q",
        "q": "r",
        "r": "s",
        "s": "t",
        "t": "u",
        "u": "v",
        "v": "w",
        "w": "x",
        "x": "y",
        "y": "z",
        "z": "a",
    }

    encrypt_map2 = {
        "a": "z",
        "b": "y",
        "c": "x",
        "d": "w",
        "e": "v",
        "f": "u",
        "g": "t",
        "h": "s",
        "i": "r",
        "j": "q",
        "k": "p",
        "l": "o",
        "m": "n",
        "n": "m",
        "o": "l",
        "p": "k",
        "q": "j",
        "r": "i",
        "s": "h",
        "t": "g",
        "u": "f",
        "v": "e",
        "w": "d",
        "x": "c",
        "y": "b",
        "z": "a",
    }

    repeat_pattern = "12"

    encrypted_string = polyalphabetic_encrypt("hello", encrypt_map1, encrypt_map2, repeat_pattern)

    print(encrypted_string)
