<!--
Copyright 2012 The Authors. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.
-->

grammar
=======

Random sentence generator based on a context-free grammar.

grammar randomly generates sentences based on a user-supplied context-free grammar.

Format
======

The definition consists of a series of newline-separated "rules."  Each rule corresponds to a part of speech, and follows the syntax &lt;part&gt; &lt;definition&gt;.  For example:

<table>
  <tr>
    <td>NP</td><td>Adj Noun</td>
  </tr>
</table>

In this rule, a noun phrase ("NP") is defined as consisting of an adjective followed by a noun.  Any part of speech which is not later defined by a rule is considered a word on its own.  For example:

<table>
  <tr>
    <td>NP</td><td>Adj Noun</td>
  </tr>
  <tr>
    <td>Adj</td><td>blue</td>
  </tr>
  <tr>
    <td>Adj</td><td>red</td>
  </tr>
</table>

This set of rules will defines "Noun" as its own word since it is nowhere defined.  However, since "Adj" is defined, it will resolve to either "blue" or "red," but never the literal "Adj."  (Note that definitions are processed serially, and thus parts of speech must be seen before they are defined)  This example illustrates another feature of rules, which is that multiple rules are allowed, and will be chosen among randomly when resolving a part of speech.

Finally, the first line of the definition does not follow the normal syntax, but instead simply defines the structure of a whole sentence.  For example:

<table>
  <tr>
    <td>Prep NP VP</td><td></td>
  </tr>
  <tr>
    <td>NP</td><td>Adj Noun</td>
  </tr>
  <tr>
    <td>VP</td><td>Adv Verb</td>
  </tr>
  <tr>
    <td>Prep</td><td>a</td>
  </tr>
  <tr>
    <td>Prep</td><td>the</td>
  </tr>
  <tr>
    <td>Adj</td><td>blue</td>
  </tr>
  <tr>
    <td>Adj</td><td>red</td>
  </tr>
  <tr>
    <td>Noun</td><td>bird</td>
  </tr>
  <tr>
    <td>Noun</td><td>dog</td>
  </tr>
  <tr>
    <td>Adv</td><td>quickly</td>
  </tr>
  <tr>
    <td>Adv</td><td>slowly</td>
  </tr>
  <tr>
    <td>Verb</td><td>ran</td>
  </tr>
  <tr>
    <td>Verb</td><td>flew</td>
  </tr>
</table>

This set of rules will create sentences such as "the red dog quickly flew."

Of course, grammar definitions are as flexible as any context-free grammar.  Consider adding features such as plural vs singular phrases, or make up your own nonsensical definitions.
