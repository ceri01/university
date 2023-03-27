package tema_2_algebretta;

/*	Questa classe estende la classe MatriceDensa, e rappresenta una matrice identita', ovvero una 
 *	matrice formata da 0 tranne che nella diagonale principale, in questa saranno presenti solo
 *	degli 1.
 *
 *	MatriceDensa e' immutabile, di conseguenza anche questa lo e'.
 *
 * 	Il costruttore di questa classe accetta solo la dimensione perche' sarebbe inutile far creare all'utente
 * 	n vettori formati da tutti 0 e un 1, per creare una matrice identita', cosi basta inserire la dimensione 
 * 	e verra' fatto in automatico.	
 */

public class MatriceIdentita extends MatriceDensa {
	// Costruttore
	
	/*	Costruttore di MatriceIdentita
	 *
	 *	pre-condizioni: l non deve contenere vettori i cui valori siano diversi da 0 e da 1, puo' esserci
	 *					al massimom un 1 per vettore e in piu' l'uno di seve trovare nella posizione che
	 *					corrispondera' poi alla diagonale della matrice identita'
	 *	post-condizioni: viene creato un oggetto istanza di MatriceIdentita
	 *
	 */
	public MatriceIdentita(int dim) {
		if(dim <= 0) throw new IllegalArgumentException("Non e' possibile creare una matrice identita' di dimensione <= 0");
		List<Vettore> v = new ArrayList<Vettore>();
		for(int i = 0; i < dim; i++) {
			int[] arr = new int[dim];
			arr[i] = 1;
			v.add(arr);
		}
		super(v);
	}

    /*  Costruttore copia di MatriceIdentita
     *
     *  per-condizioni: other != null
     *  post-condizioni: crea un oggetto che e' la copia di other
     */

	public MatriceNulla(MatriceNulla other) {
		super(other);
	}	

	/*	il prodotto di una matrice identita' per un'altra matrice ritorna la martice non identita' coinvolta
	 *	nel prodotto, infatti la matrice identita' e' l'elemeto neutro del prodotto.
	 */       
	
	@Override
    public Matrice per_matrice(final Matrice m) {
        MatriceDensa tmp = new MatriceDensa(m);
		return tmp;
    }
	
	/*	Il prodotto tra una matrice identita' e un vettore ha come risultato il vettore con cui la matrice
	 *	e' stata moltiplicata, quindi non e' necessario rieseguire il calcolo.
	 */

	@Override
	public Vettore per_vettore(final Vettore v) {
		Vettore ris = new VettoreDenso(Vettore v);
		return ris;
	}
}
