import java.util.ArrayList;
import java.util.Arrays;

public class Advent01 {
    public static Double[] weightsInput = {130762d, 108691d, 131618d, 138163d, 59967d, 130453d, 117515d, 115776d, 134083d, 86966d, 128075d, 55569d, 112843d, 97878d, 92330d, 70917d, 143903d, 81171d, 148506d, 141379d, 131161d, 88719d, 69654d, 82141d, 55265d, 75623d, 97408d, 105269d, 147378d, 126054d, 133962d, 60304d, 130503d, 138350d, 93164d, 69661d, 69271d, 100054d, 138295d, 142865d, 64142d, 123466d, 80101d, 149696d, 102510d, 129988d, 87742d, 106785d, 133039d, 59192d, 86544d, 124950d, 64242d, 80128d, 109287d, 129634d, 140335d, 118220d, 106819d, 97296d, 111003d, 103222d, 54192d, 103548d, 63861d, 140571d, 50476d, 100570d, 114065d, 110279d, 64720d, 91941d, 62312d, 80834d, 132969d, 51973d, 115887d, 68662d, 138266d, 107234d, 75795d, 81409d, 78610d, 112587d, 92384d, 111804d, 138861d, 79393d, 81285d, 131307d, 68815d, 54976d, 127529d, 103359d, 138537d, 79663d, 128097d, 56085d, 96504d, 119501d};
    public static double fuel = 3399394;

    public static void main(String[] args) {
        int result = 0;
        ArrayList<Double> weights = new ArrayList(Arrays.asList(weightsInput));
        for (Double weight : weights) {
            result += (Math.floor(weight / 3) - 2) + addFuelForFuel(Math.floor(weight / 3) - 2);
        }
        ;
        System.out.println(result);
    }

    static Double addFuelForFuel(Double fuelForModul) {
        Double fuelForFuel = 0d;
        Double zs = fuelForModul;
        while ((Math.floor(zs / 3) - 2) > 0) {
            zs = Math.floor(zs / 3) - 2;
            fuelForFuel += zs;
        }
        return fuelForFuel;
    }

}
//Resultat 1: 5099048
//Resultat 2: 5096223