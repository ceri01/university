package filesystem;

/*	La classe File e' un implementazione dell'interfaccia Entry e rappresenta un file presente nel filesystem.
 *	File e' immutabile in quanto non e' richiesto che un file possa cambiare nome, dimensione o percorso.
 *	
 *	i metodi di questa classe sono descritti nell'interfaccia Entry
 *
 *	ABS FUN: ABS(name, size, path) = Un file viene rappresentato tramite una stringa (name) che rappresenta il
 *									 nome del file, un intero (size) che rappresenta la dimensione del file e infine
 *									 un attributo di tipo Path (path) che rappresenta il percorso del file
 *
 *	REP INV: name != null
 *			 size != null && size >= 0
 *			 path != null 
 */

public class File implements Entry {
	
	// Attributi		
	private final String name;
	private int size;
	private final Path path
	
	// Costruttore
	public File(final int size, Path path) {
		if(size == null && size < 0) throw new IllegalArgumentException("La dimensione del file non puo' essere nulla o minore di 0");
		if(path == null) throw new IllegalArgumentException("Il percorso del file non puo' essere nullo");
		this.name = path.name();	
		this.size = size;
		this.path = path;
	}

	// Metodi
    
	public String name() {
        String name = new String(this.path.name());
        return name;
    }   

    public int size() {
        int size = this.size;
        return size;
    }   

    public Path path() {
        Path p = new Path(this.path);
        return p;
    }   
	

}
