package tech.sourced;

import com.ansorgit.plugins.bash.lang.BashVersion;
import org.apache.commons.lang.StringEscapeUtils;

public class Main {
    public static void main  (String args[]) {
        String code;
        code = "#!/bin/bash\necho 3; echo ${a}";
        code = "echo 3\necho ${a}";
        //code = "#!/bin/bash\n";
        System.out.println("CODE:\n" + StringEscapeUtils.escapeJava(code.toString()) + "\n\n");

        BashVersion version = BashVersion.Bash_v4;
        System.out.println("BASH VERSION:" + version + "\n\n");
        System.out.flush();

        LexerDemo.run(code, version);
        System.out.flush();

        ParserDemo.run(code, version);
        System.out.flush();
    }
}
