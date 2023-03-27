import java.util.ArrayList;
import java.util.List;
import java.util.Map;
import java.util.TreeMap;
import java.util.Comparator;
import java.util.Collections;

public class MapExampleComparator {
    private record Persona(String nome, int eta){}

    public static void main(String[] args) {
        // lista di persone
        // List of e' un metodo statico che crea una lista a
        // partire da un elenco di cosepassato
        List<Persona> persone = List.of(
                new Persona("Aldo", 32),
                new Persona("Alex", 4),
                new Persona("Giulia", 22),
                new Persona("Guido", 67),
                new Persona("Aron", 4),
                new Persona("Anna", 67),
                new Persona("Sara", 22),
                new Persona("Ciro", 22)
        );
        // Mappa che mappa ciascun eta con l'insieme delle persone di quell'eta'
        Map<Integer, List<Persona>> eta2persone = new TreeMap<>();
        // riempieamo la lista
        for(final Persona p : persone) {
            if(eta2persone.containsKey(p.eta)) {
                eta2persone.get(p.eta).add(p);
            } else {
                final List<Persona> coetanei = new ArrayList<>();
                coetanei.add(p);
                eta2persone.put(p.eta, coetanei);
            }
        }

        List<Integer> tutteEta = new ArrayList<>(eta2persone.keySet()); // creo una lista dove metto le chiavi della mappa eta2persone
        Collections.sort(tutteEta); // ordino le chiavi

        for(final Integer eta: tutteEta) {

            final List<Persona> coetanei = eta2persone.get(eta); // data un eta prendo la lista di persone associata a quell'eta'. la chiamo coetanei

            coetanei.sort(new Comparator<>() {
                @Override
                public int compare(Persona o1, Persona o2) {
                    // return o1.nome.compareTo(o2.nome); ordine crescente
                    // return -(o1.nome.compareTo(o2.nome)); ordine decrescente
                    return Integer.compare(o1.nome.length(), o2.nome.length()); // ordine per lunghezza del nome
                }
            });

            System.out.println(coetanei.size() + " persone di eta: " + eta); // stampo eta

            for(final Persona p : coetanei) { // stampo i nomi sortati dei vari coetanei
                System.out.println("\t" + p.nome);
            }
        }
    }
}


