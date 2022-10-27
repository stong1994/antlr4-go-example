// Generated from java-escape by ANTLR 4.11.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link arrayParser}.
 */
public interface arrayListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link arrayParser#vec4}.
	 * @param ctx the parse tree
	 */
	void enterVec4(arrayParser.Vec4Context ctx);
	/**
	 * Exit a parse tree produced by {@link arrayParser#vec4}.
	 * @param ctx the parse tree
	 */
	void exitVec4(arrayParser.Vec4Context ctx);
	/**
	 * Enter a parse tree produced by {@link arrayParser#ints}.
	 * @param ctx the parse tree
	 */
	void enterInts(arrayParser.IntsContext ctx);
	/**
	 * Exit a parse tree produced by {@link arrayParser#ints}.
	 * @param ctx the parse tree
	 */
	void exitInts(arrayParser.IntsContext ctx);
}