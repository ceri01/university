/**
 * Eccezione incontrollata perche dipende da un errore del programmatore, o comunque un errore di che non dipende da cause esterne (esempio: la rete)
 */
public class NegativeDegreeException extends RuntimeException {
    
    public NegativeDegreeException() {
        super();
    }

    public NegativeDegreeException(String message) {
        super(message);
    }
}
