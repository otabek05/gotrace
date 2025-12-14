/**
 * Checks if a port is valid and within the specified range.
 * @param port - The port to check (string or number)
 * @param min - Minimum allowed port (inclusive)
 * @param max - Maximum allowed port (inclusive)
 */
 export function validatePort(port: string | number, min = 1, max = 65535) {
  // Convert to number
  const portNumber = Number(port);
  if (isNaN(portNumber) || !Number.isInteger(portNumber)) {
    alert(`Port "${port}" is not a valid number.`);
    return false;
  }

  if (portNumber < min || portNumber > max) {
    alert(`Port "${portNumber}" is out of range (${min}-${max}).`);
    return false;
  }

  return true; // valid port
}
