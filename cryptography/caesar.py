

def offset_encrypt(input: str, offset: int) -> str:
    """
    Encrypts a string by shifting each character by the given offset.

    Args:
        input: The string to encrypt.
        offset: The amount to shift each character by.

    Returns:
        The encrypted string.
    """
    encrypted_string = ""
    for character in input:
        encrypted_string += chr(ord(character) + offset)

    return encrypted_string
    