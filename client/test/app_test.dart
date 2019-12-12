@TestOn('browser')

import 'package:test/test.dart';

void main() {
    test('default', () {
        print("test example.");
    });

    // final testBed =
    // NgTestBed.forComponent<AppComponent>(ng.AppComponentNgFactory);
    // NgTestFixture<AppComponent> fixture;

    // setUp(() async {
    //     fixture = await testBed.create();
    // });

    // tearDown(disposeAnyRunningTest);

    // test('Default greeting', () {
    //     expect(fixture.text, 'Angular');
    // });

    // test('Greet world', () async {
    //     await fixture.update((c) => c.title = 'World');
    //     expect(fixture.text, 'Hello World');
    // });

    // test('Greet world HTML', () {
    //     final html = fixture.rootElement.innerHtml;
    //     expect(html, '<h1>Hello Angular</h1>');
    // });
}
