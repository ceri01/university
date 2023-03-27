package tema_2_algebretta;

/*	Questa classe estende la classe MatriceDensa, e rappresenta una matrice nulla, ovvero una matrice formata
 *	solo da 0.
 *
 *	MatriceDensa e' immutabile, di conseguenza anche questa lo e'.
 *
 * 	Il costruttore di questa classe accetta solo la dimensione perche' sarebbe inutile far creare all'utente
 * 	n vettori da n zeri per creare una matrice nulla, cosi basta inserire la dimensione e verra' fatto in
 * 	automatico	
 */

public class MatriceNulla extends MatriceDensa {
	// Costruttore
	
	/*	Costruttore di MatriceNulla
	 *
	 *	pre-condizioni: l non deve contenere vettori i cui valori siano diversi da 0
	 *	post-condizioni: viene creato un oggetto istanza di MatriceNulla
	 *
	 */
	public MatriceNulla(int dim) {
		if(dim <= 0) throw new IllegalArgumentException("Non e' possibile creare una matrice nulla di dimens    ione <= 0");
		List<Vettore> v = new ArrayList<Vettore>();
		for(int i = 0; i < dim; i++) {
			int[] arr = new int[dim]; // di default array gia inizializzato a 0
			v.add(arr);
		}
		super(v);
	}

    /*  Costruttore copia di MatriceNulla
     *
     *  per-condizioni: other != null
     *  post-condizioni: crea un oggetto che e' la copia di other
     */

	public MatriceNulla(MatriceNulla other) {
		super(other);
	}

	// Metodi

	/*	La somma di una matrice per una matrice nulla ha come risultato la matrice non nulla coinvolta
	 *	nella somma, quindi non serve eseguire il calcolo
	 */
	
	@Override
    public Matrice piu(final Matrice m) {
        if(m.dim() != this.dim) throw new IllegalArgumentException("Le matrici devono avere la stessa dimensione per essere sommate.");
		MatriceDensa tmp = new MatriceDensa(m);
		return tmp;
    }       
	
	/*	Il prodotto di una matrice nulla per uno scalare ritorna una matrice nulla, quindi non c'e'
	 *	bisogno di eseguire il calcolo.
	 */

	@Override
    public Matrice per_scalare(final int i) {
		return MatriceNulla(this);
    }

	/*	Il prodotto di una matrice nulla per un'altra matrice ritorna una matrice  nulla, quindi non c'e'
	 *	bisogno di eseguire il calcolo.
	 */       
	
	@Override
    public Matrice per_matrice(final Matrice m) {
        if(m.dim() != this.dim) throw new IllegalArgumentException("Le matrici devono avere la stessa dimensione per essere moltiplicate.");
		return MatriceNulla(this);
    }
	
	/*	Il prodotto tra una matrice nulla e un vettore da come risultato un vettore nullo, quindi non c'e'
	 *	bisogno di eseguire il calcolo.
	 */

	@Override
	public Vettore per_vettore(final Vettore v) {	
        if(v.dim() != this.dim) throw new IllegalArgumentException("Il vettore deve avere la stessa dimensione della matrice per poter eseguire il prodotto.");
		Vettore v = new VettoreDenso(v.dim());
		return v;
	}	
}
