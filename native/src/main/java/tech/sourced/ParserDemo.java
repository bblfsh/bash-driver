package tech.sourced;

import com.ansorgit.plugins.bash.lang.BashVersion;
import com.ansorgit.plugins.bash.lang.lexer.BashLexer;
import com.ansorgit.plugins.bash.lang.parser.BashParserDefinition;
import com.ansorgit.plugins.bash.lang.parser.BashPsiBuilder;
import com.ansorgit.plugins.bash.lang.parser.FileParsing;
import com.intellij.core.CoreApplicationEnvironment;
import com.intellij.core.CoreProjectEnvironment;
import com.intellij.lang.ASTNode;
import com.intellij.lang.PsiBuilder;
import com.intellij.lang.impl.PsiBuilderFactoryImpl;
import com.intellij.mock.MockProject;
import com.intellij.openapi.Disposable;

public class ParserDemo {
    public static void run(CharSequence code, BashVersion version) {
        BashPsiBuilder b = builder(code, version);
        FileParsing fp = new FileParsing();

        fp.parseFile(b);

        System.out.println(b.getOriginalText());
        System.out.flush();

        ASTNode root = b.getTreeBuilt();
        report(root);
        System.out.flush();
    }

    private static BashPsiBuilder builder(CharSequence code, BashVersion version) {
        com.intellij.openapi.extensions.ExtensionsArea rootArea = com.intellij.openapi.extensions.Extensions.getRootArea();
        CoreApplicationEnvironment.registerExtensionPoint(rootArea, com.intellij.lang.MetaLanguage.EP_NAME, com.intellij.lang.MetaLanguage.class);

        CoreApplicationEnvironment appEnv = new CoreApplicationEnvironment(new Disposable() {
            @Override
            public void dispose() {
            }});

        CoreProjectEnvironment environment = new CoreProjectEnvironment(new Disposable() {
            @Override
            public void dispose() {
            }
        }, appEnv);

        MockProject project = environment.getProject();
        PsiBuilder builder = new PsiBuilderFactoryImpl().createBuilder(
                new BashParserDefinition(),
                new BashLexer(),
                code
        );
        return new BashPsiBuilder(project, builder, version);
    }

    private static void report(ASTNode root) {
        System.out.println(root);
    }
}
