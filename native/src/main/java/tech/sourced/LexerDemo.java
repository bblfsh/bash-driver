package tech.sourced;

import com.ansorgit.plugins.bash.lang.BashVersion;
import com.ansorgit.plugins.bash.lang.lexer.BashLexer;
import com.intellij.psi.tree.IElementType;

import java.util.ArrayList;
import java.util.List;

public class LexerDemo {
    public static void run(CharSequence code, BashVersion version) {
        List<Token> tokens = lex(code, version);

        System.out.println("TOKENS:");
        for (Token t : tokens) {
            System.out.println(t);
        }
        System.out.println("\n\n");
    }

    private static List<Token> lex(CharSequence code, BashVersion version) {
        List<Token> l = new ArrayList<Token>();

        BashLexer lexer = new BashLexer(version);

        lexer.start(code);
        while (true) {
            IElementType type = lexer.getTokenType();
            if (type == null) {
                break;
            }
            l.add(Token.from(lexer));
            lexer.advance();
        }

        return l;
    }
}