

export function convertImagePathToLink(protocol, host, imagePath) {
    if (imagePath) {
        return `${protocol}://${host}${imagePath}`;
    }
    return null;
}