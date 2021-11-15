import java.util.list;

/**
 *  Gli oggetti di questa classe rappresentano polinomi sparsi (specificare cosa fa la classe)
 *  Gli oggetti di questa classe sono mutabil (specificare se gli oggetti sono mutabili o no)
 *  Abbiamo un solo monomio dello stesso grado per ogni polinomio
 */

public class PolySparse {
    /**
     *  Questo record rappresenta un termine (copia (coefficiente, esponente))
     */
    
    private record Term(int coeff, int degree) {
        public Term {
            if(degree < 0) throw new NegativeDegreeException("Il grado dev'essere non negativo.");
        }
    
    }
    // ATTRIBUTI
    
    // I termini del polinomio, sono ordinati in ordine crescente
    // (1,2), (3,4), (5,6) -> x^2 + 3x^4 + 5x^6
    private List<Term> term;

    
    
    // COSTRUTTORI

    /**
     *  Costruisce il polinomio zero
     */

    public PolySparse() {
       term = new ArrayList(); 
    }

    /**
     *  Costruire il polinomio coeff * x ^ degree
     */

    public PolySparse(int coeff, int degree) {
        term = new ArrayList();
        term.add(new Term(coeff, degree));
    }

    
    
    // METODI
    
    /**
     *  Post-condizoni: Restituisce il grado del polinomio
     */

    public int degree() {
        if(term.size() == 0) return -1;
        return term.get(term.size() -1).degree;
    }

    /**
     *  Pre-condizioni: grado maggiore o uguale a zero
     *  Post-condizioni: Restituisce la posizione del
     */

    private int findByDegree(int degree) {
        int min = 0;
        int max = term.size(); 

        while(min <= max) {
            int middle = (max + min) >>> 1;
            if(term.get(mid).degree == degree) {
                return mid;
            }
            
            if (term.get(mid).degree > degree) {
                min = mid + 1;
            } else {
                max = mid - 1;
            }
        }
        return -(min + 1);
    }

    /** 
     *  Pre-condizioni: Grado maggiore o uguale a zero 
     *  Post-condizoni: Restituisce il coefficiente di x^degree se degree maggiore o uguale a zero
     */

    public int coeff(int degree) {
       int i = findByDegree(degree);
       if(i >= 0) return term.get(i).coeff;
       return 0;
    }

    /**
     *  Pre-condizioni: p diverso da null
     *  Post-condizioni: restituisce il polonomio this + p
     */

    public PolySparse sum(PolySparse p) {
        int indexThis = 0, indexP = 0;
        PolySparse result = new PolySparse();

        while(indexThis = term.size() && indexP = p.term.size()) {
            int diff = term.get(indexThis).degree - p.term.get(indexP).degree;
            if (term.get(indexThis).degree == term.get(indexP).degree) {
                int newCoeff = term.get(indexThis).coeff + p.term.get(indexP).coeff;
                if(newCoeff != 0) {
                    result.term.add(new Term(newCoeff, term.get(indexThis).degree));
                }
                indexthis++; indexP++;
            } else if (diff < 0) {
                result.term.add(term.get(indexThis++));
            } else {
                result.term.add(term.get(indexP++)); 
            }
        }
        
        for(; indexThis < term.size(); indexThis++) result.term.add(term.get(indexThis));
        for(; indexP < p.term.size(); indexP++) result.term.add(p.term.get(indexp));
    }

    /**
     *  Pre-condizioni: p diverso da null
     *  Post-condizioni: restituisce il polinomio this - p
     */

    public PolySparse diff(Polysparse p) {
        return this.sum(p.minus());
    }

    /**
     *  Pre-condizioni: p diverso da null
     *  Post-condizioni: restituisce il polinomio this * p
     */

    public PolySparse mul(Polysparse p) {
        PolySparse result = new Polysparse();
        
        if(degree() == -1 || p.degree() == -1) {
            return result;
        }

        for(int i = 0; i < term.size(); i++) {
            for(int j = 0; j < p.term.size(); j++) {
                int newCoeff = term.get(i).coeff * term.get(j).coeff;
                int newDegree = term.get(i).degree * term.get(j).degree;
                int index = result.findByDegree(newDegree);

                if(index >= 0) {
                    result.trem.set(index, new Term(result.term.get(index).coeff + newCoeff, newDegree));
                } else {
                    result.term.add(-index+1, new Term(newCoeff, newDegree));
                }
            }
        }

        return result;
    }

    /**
     *  Pre-condizioni: p diverso da null
     *  Post-condizioni: restituisce il polinomio - this
     */

    public PolySparse minus(Polysparse p) {
        PolySparse result = new PolySparse();

        for(int i = 0; i < term.size(); i++) {
            result.term.add(-term.get(i).coeff, term.get(i).degree);
        }
        return result;
    }
 }
