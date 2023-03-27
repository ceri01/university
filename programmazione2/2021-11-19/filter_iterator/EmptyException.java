/**
 *  Eccezione non controllata, viene chiamata quando si tenta di accedere ad un elemento dell'insieme vuoto
 */

class EmptyException extends RuntimeException {
    public EmptyException() {
        super();
    }

    public EmptyException(String message) {
        super(message);
    }
}
