

def monoalphabetic_encrypt(input: str, encrypt_map: dict) -> str:
    """
    Encrypts a string using a monoalphabetic cipher.

    Args:
        input: The string to encrypt.
        encrypt_map: The mapping of characters to encrypted characters.

    Returns:
        The encrypted string.
    """

    encrypted_string = ""
    for character in input:
        encrypted_character = encrypt_map.get(character)
        if encrypted_character is None:
            encrypted_character = character
        encrypted_string += encrypted_character

    return encrypted_string

if __name__ == "__main__":
    encrypt_map = {
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
    
    plaintext = "How are you?"
    print(monoalphabetic_encrypt(plaintext, encrypt_map))