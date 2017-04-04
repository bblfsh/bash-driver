package tech.sourced;

import com.intellij.lexer.Lexer;
import com.intellij.psi.tree.IElementType;

public class Token {
    private int start;
    private int end;
    private String text;
    private IElementType type;

    // BashElementType adds this string as a prefix to the toString(), and we don't want it
    static private int prefixLen = "[Bash] ".length();

    private Token() {}

    private Token(int start, int end, String text, IElementType type) {
        this.start = start;
        this.end = end;
        this.text = text;
        this.type = type;
    }

    static Token from(Lexer lexer) {
        IElementType type = lexer.getTokenType();
        if (type == null) {
            return null;
        }
        return new Token(
                lexer.getTokenStart(),
                lexer.getTokenEnd(),
                lexer.getTokenText(),
                type
        );
    }

    public String toString() {
        return  "start=" + start + ", " +
                "end=" + end + ", " +
                "type=\"" + type.toString().substring(prefixLen) + "\", " +
                "value=\"" + org.apache.commons.lang.StringEscapeUtils.escapeJava(text) + "\"";
    }
}
