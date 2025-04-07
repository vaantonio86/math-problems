import java.util.Random;

public class GoGame {
    public static void main(String[] args) {
        Random random = new Random();
        int num1 = random.nextInt(10);
        int num2 = random.nextInt(10);
        if (num1 > 5 && num2 > 5) {
            System.out.println("You win!");
        } else {
            System.out.println("Try again.");
        }
    }
}
