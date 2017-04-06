package bblfsh.bash;

import com.ansorgit.plugins.bash.lang.BashVersion;
import com.intellij.lang.ASTNode;
import org.junit.Test;
import static org.junit.Assert.assertNotNull;

public class ViabilityTest {
    @Test
    public void canParse() {
        String code = "#!/bin/bash\na=3; echo ${a}";
        BashVersion version = BashVersion.Bash_v4;
        ASTNode root = ParserDemo.run(code, version);
        assertNotNull(root);
    }
}
