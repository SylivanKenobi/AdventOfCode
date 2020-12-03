import java.util.HashMap;

public class Advent04 {
    public static void main(String[] args) {
        int maxValue = 767346;
        int minValue = 231832;
        int counter = 0;
        for (int i = minValue; i <= maxValue; i++) {
            if (checkAdjacentNumber(i + "")) {
                if (checkIncrease(i + "")) {
                    counter++;
                }
            }

        }
        System.out.println(counter);
    }

    static boolean checkIncrease(String pin) {
        for (int i = 0; i < pin.length() - 1; i++) {
            if (pin.charAt(i) > pin.charAt(i + 1)) return false;
        }
        return true;
    }

    static boolean checkAdjacentNumber(String pin) {
        HashMap<Character, Integer> fuck = new HashMap<>();
        for (int i = 0; i < pin.length() - 1; i++) {
            if (pin.charAt(i) == pin.charAt(i + 1)) {
                fuck.put(pin.charAt(i), fuck.containsKey(pin.charAt(i)) ? fuck.get(pin.charAt(i)) + 1 : 2);
            }
        }
        return fuck.containsValue(2);
    }
}
