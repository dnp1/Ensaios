import 'package:angular2/core.dart';
import 'package:angular2_tour_of_heroes/hero_service.dart';
import 'package:angular2_tour_of_heroes/heroes_component.dart';
import 'package:angular2/router.dart';
import 'package:angular2_tour_of_heroes/dashboard_component.dart';
import 'package:angular2_tour_of_heroes/hero_detail_component.dart';

@Component(
    selector: 'my-app',
    template: '''
    <h1>{{title}}</h1>
    <nav>
    <a [routerLink]="['Dashboard']">Dashboard</a>
    <a [routerLink]="['Heroes']">Heroes</a>
    </nav>
    <router-outlet></router-outlet>
''',
    directives: const [
        HeroesComponent,
        ROUTER_DIRECTIVES,
    ],
    providers: const [HeroService, ROUTER_PROVIDERS],
    styleUrls: const ['app_component.css']
)
@RouteConfig(const [
    const Route(
        path: '/heroes',
        name: 'Heroes',
        component: HeroesComponent
    ),
    const Route(
        path: '/dashboard',
        name: 'Dashboard',
        component: DashboardComponent,
        useAsDefault: true
    ),
    const Route(
        path: '/detail/:id',
        name: 'HeroDetail',
        component: HeroDetailComponent
    )
]
)
class AppComponent {
    String title = 'Tour of Heroes';
}