package bblfsh.bash;

import com.ansorgit.plugins.bash.lang.BashVersion;
import com.ansorgit.plugins.bash.lang.lexer.BashElementType;
import com.ansorgit.plugins.bash.lang.lexer.BashLexer;
import com.ansorgit.plugins.bash.lang.parser.BashParser;
import com.ansorgit.plugins.bash.lang.parser.BashParserDefinition;
import com.ansorgit.plugins.bash.lang.parser.BashPsiBuilder;
import com.ansorgit.plugins.bash.lang.parser.FileParsing;
import com.intellij.core.CoreApplicationEnvironment;
import com.intellij.core.CoreProjectEnvironment;
import com.intellij.lang.ASTNode;
import com.intellij.lang.ParserDefinition;
import com.intellij.lang.PsiBuilder;
import com.intellij.lang.impl.PsiBuilderFactoryImpl;
import com.intellij.mock.MockProject;
import com.intellij.openapi.Disposable;
import com.intellij.psi.tree.IElementType;
import com.intellij.openapi.extensions.ExtensionsArea;
import com.intellij.lang.MetaLanguage;

import org.apache.commons.lang.StringEscapeUtils;
import org.apache.commons.lang.StringUtils;

import java.util.Properties;

public class ParserDemo {
    public static ASTNode run(CharSequence code, BashVersion version) {
        MockProject project = project();
        BashParser parser = new BashParser(project, version);
        ParserDefinition parserDefinition = new BashParserDefinition();

        PsiBuilder builder = builder(parserDefinition, code);
        IElementType root = parserDefinition.getFileNodeType();

        return parser.parse(root, builder);
    }

    private static MockProject project() {
        ExtensionsArea rootArea = Extensions.getRootArea();
        CoreApplicationEnvironment.registerExtensionPoint(rootArea, MetaLanguage.EP_NAME, MetaLanguage.class);

        CoreApplicationEnvironment appEnv = new CoreApplicationEnvironment(new Disposable() {
            @Override
            public void dispose() {
            }});

        CoreProjectEnvironment environment = new CoreProjectEnvironment(new Disposable() {
            @Override
            public void dispose() {
            }
        }, appEnv);

        return environment.getProject();
    }

    private static PsiBuilder builder(ParserDefinition parserDefinition, CharSequence code) {
        return new PsiBuilderFactoryImpl().createBuilder(
                parserDefinition,
                new BashLexer(),
                code
        );
    }
}
