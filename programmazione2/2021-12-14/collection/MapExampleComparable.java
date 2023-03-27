import java.util.ArrayList;
import java.util.List;
import java.util.Map;
import java.util.TreeMap;
import java.util.Collections;


public class MapExampleComparable {
    private record Persona(String nome, int eta) implements Comparable<Persona> {
        // Implementiamo comparable per poter comparare tutti gli oggetti di tipo persona
        @Override
        public int compareTo(Persona o) {
            return nome.compareTo(o.nome);
        }
        
        // obbigatorio implementare anche equals
        @Override
        public boolean equals(Object o) {
            if (!(o instanceof Persona)) {
                return false;
            }
            return ((Persona)o).nome.equals(nome);
        }
        // di conseguenza anche hashcode
        @Override
        public int hashCode() {
            return nome.hashCode();
        }
    }

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

            /* Se non avessi implementato compareTo per persona avrei dovuto fare cosi
            final List<String> nomi = new ArrayList<>(); // creo una lista di nomi che conterra' i nomi di tutti i coetanei
            for(final Persona p : coetanei) { // popolo nomi con tutti i nomi dei coetanei
                nomi.add(p.nome);
            }

            // ho divuto fare questa lista di nomi perche' le stringhe sono confrontabili e posso eseguire il sort.
            // se applico il sort alla lista di Persona ho un errore perche gli elementi di tipo Persona non sono confrontabili
            Collections.sort(nomi); // utilizzo il metodo sort della classe collections per sortare i nomi
            */
            Collections.sort(coetanei);
            System.out.println(coetanei.size() + " persone di eta: " + eta); // stampo eta

            for(final Persona p : coetanei) { // stampo i nomi sortati dei vari coetanei
                System.out.println("\t" + p.nome);
            }
        }
    }
}


